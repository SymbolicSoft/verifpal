# This file was generated by GoReleaser. DO NOT EDIT.
class Verifpal < Formula
  desc "Cryptographic protocol analysis for students and engineers."
  homepage "https://verifpal.com"
  version "0.15.6"
  bottle :unneeded

  if OS.mac?
    url "https://source.symbolic.software/verifpal/verifpal/uploads/be3904adc694f8e20d708be9f6be9c56/verifpal_0.15.6_macos_amd64.zip"
    sha256 "0bffbe6d307a3ed696fedbbabbb7d6f233ef44b9d68cc7179391686f6e98af86"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://source.symbolic.software/verifpal/verifpal/uploads/73d303a40dba4680fcfc340b0820fc9c/verifpal_0.15.6_linux_amd64.zip"
      sha256 "49d49378e952069a1292956fe0f0ef5f2d9df2369a1d575f2042f2a4f7b8b7b4"
    end
  end

  def install
    bin.install "verifpal"
  end
end
