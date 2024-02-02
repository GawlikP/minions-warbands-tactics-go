package texture

import (
  "log"
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/text"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "golang.org/x/image/font"
  "minions-warbands-tactics/constant"
  "fmt"
)

func DrawCenteredText(screen *ebiten.Image, font font.Face, s string, cx, cy int) {
    bounds := text.BoundString(font, s)
    x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
    text.Draw(screen, s, font, x, y,  color.White)
}

func LoadTexture(path string) *ebiten.Image {
  img, _, err := ebitenutil.NewImageFromFile(path)
  if err != nil {
    log.Fatal(err)
  }
  return img
}

func FlipHorizontal(source *ebiten.Image) *ebiten.Image {
    result := ebiten.NewImage(source.Bounds().Max.X, source.Bounds().Max.Y)
    op := ebiten.DrawImageOptions{}
    op.GeoM.Scale(-1, 1)
    op.GeoM.Translate(float64(source.Bounds().Dy()), 0)
    result.DrawImage(source, &op)
    return result
}

func ScaleTexture(texture *ebiten.Image, width int, height int) *ebiten.Image {
  var scaledTexture *ebiten.Image
  var op ebiten.DrawImageOptions
  imageW, imageH := texture.Bounds().Dx(), texture.Bounds().Dy()
  subImage := ebiten.NewImage(width, height)
  scaledTexture = ebiten.NewImageFromImage(subImage)
  op.GeoM.Reset()
  op.GeoM.Scale(float64(width)/float64(imageW), float64(height)/float64(imageH))
  scaledTexture.DrawImage(texture, &op)
  return scaledTexture
}

func LoadAnimationFrames(name string, animationType string, frames int, flip bool) []*ebiten.Image {
  var img *ebiten.Image
  var array []*ebiten.Image
  for i := 0; i < frames; i++ {
    index := i + 1
    path := fmt.Sprintf("images/%s/%s%s%d.png", name, name, animationType, index)
    img = LoadTexture(path)
    if flip {
      img = FlipHorizontal(img)
    }
    array = append(array, ScaleTexture(img, constant.UNITSIZE, constant.UNITSIZE))
  }
  return array
}
