components {
  id: "player"
  component: "/res/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player-idle-down\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/res/sprites.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "rocketfactory"
  type: "factory"
  data: "prototype: \"/main/rocket.go\"\n"
  ""
}
