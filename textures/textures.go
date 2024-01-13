package textures

import (
  "log"
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
  "minions-warbands-tactics/constants"
  "io/ioutil"
)

type Tex struct {
  StandardFont          font.Face
  Cursor                *ebiten.Image
  NewGameBanner         *ebiten.Image
  ExitGameBanner        *ebiten.Image
  GrassTile             *ebiten.Image
  StoneTile             *ebiten.Image
  SandTile              *ebiten.Image
  UIBadge               *ebiten.Image
  //Minions
  FishMinion            *ebiten.Image
  RatMinion             *ebiten.Image
  BazaltieWalkingRight  []*ebiten.Image
  BazaltieWalkingLeft   []*ebiten.Image
  BazaltieWalkingUp     []*ebiten.Image
  BazaltieWalkingDown   []*ebiten.Image
  ThreedyWalkingRight   []*ebiten.Image
  ThreedyWalkingLeft    []*ebiten.Image
  ThreedyWalkingUp       []*ebiten.Image
  ThreedyWalkingDown     []*ebiten.Image
}
// TODO: REFACTOR INITIALIZATION
func (t *Tex) InitTextures() {
  var err error
  var img *ebiten.Image
  err = t.InitFonts()
  if err != nil {
    log.Fatal(err)
  }
  img = LoadTexture("images/cursor.png")
  t.Cursor = ScaleTexture(img, constants.CURSORSIZE, constants.CURSORSIZE)
  log.Print("Initialized the Cursor Texture")

  img = LoadTexture("images/ui_badge.png")
  t.UIBadge = ScaleTexture(img, constants.BANNERWIDTH, constants.BANNERHEIGHT)
  log.Print("Initialized the UIBadge Texture")


  // TILES INITIALIZATION
  log.Print("Initializing the tiles textures")

  img = LoadTexture("images/grass_tile.png")
  t.GrassTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the GrassTile Texture")

  img = LoadTexture("images/StoneWall.png")
  t.StoneTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the StoneTile Texture")

  img = LoadTexture("images/sand_tile.png")
  t.SandTile = ScaleTexture(img, constants.TILESIZE, constants.TILESIZE)
  log.Print("Initialized the SandTile Texture")

  //Units Initialization
  img = LoadTexture("images/fish_unit.png")
  t.FishMinion = ScaleTexture(img, constants.UNITSIZE, constants.UNITSIZE)
  log.Print("Initialized the Fish unit Texture")

  img = LoadTexture("images/mouse_unit.png")
  t.RatMinion = ScaleTexture(img, constants.UNITSIZE, constants.UNITSIZE)
  log.Print("Initialized the Mouse Texture")

  // PRIMITIVES
  log.Printf("Initializing the Primitives")
  t.NewGameBanner = t.GeneratePrimitiveBanner("NEW GAME", color.RGBA{200,100,100,100})
  t.ExitGameBanner = t.GeneratePrimitiveBanner("EXIT GAME", color.RGBA{200,100,100,100})
  log.Printf("Initializing the Minions Textures")
  InitMinionsTextures(t)
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
    DPI:  72,
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
  DrawCenteredText(bannerTexture, t.StandardFont, msg, constants.BANNERWIDTH/2, constants.BANNERHEIGHT/2)
  return bannerTexture
}
