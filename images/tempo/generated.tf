# DO NOT EDIT - this file is autogenerated by tfgen

output "summary" {
  value = merge(
    {
      basename(path.module) = {
        "ref"    = module.tempo.image_ref
        "config" = module.tempo.config
        "tags"   = ["latest"]
      }
  })
}

