cask "kaizen" do
  version "0.1.0"
  sha256 "93d0f64fc93d8580e96b405fb1905d43f5e98788c6c99df3976e6350637a3d32"

  url "https://github.com/sailpoint-playground/kaizen-cli/releases/download/v0.1.0/kaizen", verified: "github.com/sailpoint-playground/"
  name "kaizen"
  desc "A powerful CLI for streamlined CI/CD operations."
  homepage "https://github.com/sailpoint-playground/kaizen-cli"

  binary "kaizen"

  livecheck do
    url :homepage
    regex(/Kaizen\s+v?(\d+(?:\.\d+)+)/i)
  end
end