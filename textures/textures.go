package textures

import (
  "log"
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/text"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
  "io/ioutil"
)

type Tex struct {
  standardFont font.Face
  Cursor *ebiten.Image
  NewGameBanner *ebiten.Image
  ExitGameBanner *ebiten.Image
}

func (t *Tex) InitTextures() {
  var err error
  var img *ebiten.Image
  err = t.InitFonts()
  if err != nil {
    log.Fatal(err)
  }
  img, _, err = ebitenutil.NewImageFromFile("images/cursor.png")
  if err != nil {
    log.Fatal(err)
  }
  t.Cursor = ScaleTexture(img, 32, 32)
  log.Print("Initialized the cursor Texture")
  
  t.NewGameBanner = t.GeneratePrimitiveBanner("NEW GAME", color.RGBA{200,100,100,100})
  t.ExitGameBanner = t.GeneratePrimitiveBanner("EXIT GAME", color.RGBA{200,100,100,100})
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

func (t *Tex) InitFonts() error {
  var err error
  fontBytes, err := ioutil.ReadFile("fonts/Pixuf.ttf")
  if err != nil {
    return err
  }
  tt, err := opentype.Parse(fontBytes)
  if err != nil {
    return err
  }
  t.standardFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
    Size: 16,
    DPI:  128,
  })
  if err != nil {
    return err
  }
  return nil
}

func (t *Tex) GeneratePrimitiveBanner(msg string, c color.Color) *ebiten.Image {
  log.Printf("Generating Banner with text %s and color %v", msg, c)
  banner := ebiten.NewImage(240, 64)
  bannerTexture := ebiten.NewImageFromImage(banner)
  bannerTexture.Fill(c)
  bannerMiddle := 240/2
  textLength := len(msg)*8
  text.Draw(bannerTexture, msg, t.standardFont, bannerMiddle-textLength, 32, color.White)
  return bannerTexture
}
