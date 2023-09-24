# Content managed by Project Forge, see [projectforge.md] for details.
source = ["./build/dist/darwin_darwin_amd64_v1/npn"]
bundle_id = "dev.npn"

notarize {
  path = "./build/dist/npn_0.0.10_darwin_amd64_desktop.dmg"
  bundle_id = "dev.npn"
}

apple_id {
  username = "kyle@kyleu.com"
  password = "@env:APPLE_PASSWORD"
}

sign {
  application_identity = "Developer ID Application: Kyle Unverferth (C6S478FYLD)"
}

dmg {
  output_path = "./build/dist/npn_0.0.10_darwin_amd64.dmg"
  volume_name = "NPN"
}

zip {
  output_path = "./build/dist/npn_0.0.10_darwin_amd64_notarized.zip"
}
