# This file was generated by GoReleaser. DO NOT EDIT.
class Verifpal < Formula
  desc "Cryptographic protocol analysis for students and engineers."
  homepage "https://verifpal.com"
  version "0.19.2"
  bottle :unneeded

  if OS.mac?
    url "https://source.symbolic.software/verifpal/verifpal/uploads/691c3cd73c1c54b7476c7420630f7cfc/verifpal_0.19.2_macos_amd64.zip"
    sha256 "5d2bf41f2c3492a4777eabb4d416d5466b1e025919c6a3c76b849169a2c17fed"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://source.symbolic.software/verifpal/verifpal/uploads/1e25d3961ef510486272341b01d96846/verifpal_0.19.2_linux_amd64.zip"
      sha256 "69535ec0dc0039a7751afb47b34e9ea12d2ec4ddc6b741912805bfdec75a9805"
    end
  end

  def install
    bin.install "verifpal"
  end
end
