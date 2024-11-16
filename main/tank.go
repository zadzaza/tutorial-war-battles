components {
  id: "tank"
  component: "/res/tank.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"tank-down\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/res/sprites.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemies\"\n"
  "mask: \"rockets\"\n"
  "mask: \"enemies\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 8.988764\n"
  "  data: 10.751748\n"
  "  data: 10.0\n"
  "}\n"
  "locked_rotation: true\n"
  ""
}
