package textures

import (
  "log"
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/text"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
  "minions-warbands-tactics/constants"
  "io/ioutil"
)

type Tex struct {
  StandardFont      font.Face
  Cursor            *ebiten.Image
  NewGameBanner     *ebiten.Image
  ExitGameBanner    *ebiten.Image
  GrassTile         *ebiten.Image
  StoneTile         *ebiten.Image
  SandTile          *ebiten.Image
  UIBadge           *ebiten.Image
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
  t.Cursor = ScaleTexture(img, constants.CURSORSIZE, constants.CURSORSIZE)
  log.Print("Initialized the Cursor Texture")

  img, _, err = ebitenutil.NewImageFromFile("images/ui_badge.png")
  if err != nil {
    log.Fatal(err)
  }
  t.UIBadge = ScaleTexture(img, constants.BANNERWIDTH, constants.BANNERHEIGHT)
  log.Print("Initialized the UIBadge Texture")


  // TILES INITIALIZATION
  log.Print("Initializing the tiles textures")
  img, _, err = ebitenutil.NewImageFromFile("images/grass_tile.png")
  if err != nil {
    log.Fatal(err)
  }

  t.GrassTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the GrassTile Texture")

  img, _, err = ebitenutil.NewImageFromFile("images/stone_tile.png")
  if err != nil {
    log.Fatal(err)
  }

  t.StoneTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the StoneTile Texture")

  img, _, err = ebitenutil.NewImageFromFile("images/sand_tile.png")
  if err != nil {
    log.Fatal(err)
  }

  t.SandTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the SandTile Texture")

  // PRIMITIVES
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
  t.StandardFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
    Size: constants.FONTSIZE,
    DPI:  64,
  })
  if err != nil {
    return err
  }
  return nil
}

func (t *Tex) GeneratePrimitiveBanner(msg string, c color.Color) *ebiten.Image {
  log.Printf("Generating Banner with text %s and color %v", msg, c)
  banner := ebiten.NewImage(constants.BANNERWIDTH, constants.BANNERHEIGHT)
  bannerTexture := ebiten.NewImageFromImage(banner)
  bannerTexture.Fill(c)
  bannerMiddle := constants.BANNERWIDTH/2
  textLength := len(msg)*constants.FONTSIZE
  text.Draw(bannerTexture, msg, t.StandardFont, bannerMiddle-textLength, constants.FONTSIZE*2, color.White)
  return bannerTexture
}
