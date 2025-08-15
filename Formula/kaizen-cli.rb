# This is an auto-generated file. DO NOT EDIT
class Kaizen < Formula
    desc "Application lifecycle orchestration"
    homepage "https://kaizen.sailpoint.com/"
    version "v1.0.3"

    on_macos do
      if Hardware::CPU.intel?
        url "https://github.com/sailpoint-playground/kaizen-cli/releases/download/v1.0.3/kaizen-cli-darwin-amd64"
        sha256 "a842cb6df8da4d4bdddcc34cacdd6713481b5dd464453f90927910b24e72a219"
      elsif Hardware::CPU.arm?
        url "https://github.com/sailpoint-playground/kaizen-cli/releases/download/v1.0.3/kaizen-cli-darwin-arm64"
        sha256 "75b0876daccd5b77d8bf80344eae7f1b12fa03949186b02e64cc3b02feb89b6c"
      end
    end

    on_linux do
      if Hardware::CPU.intel?
        url "https://github.com/sailpoint-playground/kaizen-cli/releases/download/v1.0.3/kaizen-cli-linux-amd64"
        sha256 "54809d8953da99c5fee65ca7fd87a947f8914d941d0f2352c926508f697933da"
      elsif Hardware::CPU.arm?
        url "https://github.com/sailpoint-playground/kaizen-cli/releases/download/v1.0.3/kaizen-cli-linux-arm64"
        sha256 "b59deef24a56498ca25c512f59d8d107004b66a074e368737ee7fe8289883c8b"
      end
    end

    def install
      bin.install "kaizen-cli-#{OS.kernel_name.downcase}-#{Hardware::CPU.arch}" => "kaizen-cli"
    end

    test do
      assert_match version.to_s, shell_output("#{bin}/kaizen-cli --version")
    end
end
