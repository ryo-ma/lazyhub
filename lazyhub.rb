# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Lazyhub < Formula
  desc ":octocat: lazyhub - Terminal UI Client for GitHub using gocui."
  homepage "https://github.com/ryo-ma/lazyhub"
  url "https://github.com/ryo-ma/lazyhub/archive/v0.0.1.tar.gz"
  sha256 "00c442edc62c219bf81880a47dc36f0ad33d0a9dd7457b84a42e66c15224fd5d"

  # depends_on "cmake" => :build

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    # Remove unrecognized options if warned by configure
    #
    print buildpath
    ENV["GOPATH"] = buildpath
    path = buildpath/"src/github.com/ryo-ma/lazyhub/"
    path.install buildpath.children

    cd path do
      system "dep", "ensure", "-vendor-only"
      system "go", "build"
      bin.install "lazyhub"
    end
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! For Homebrew/homebrew-core
    # this will need to be a test that verifies the functionality of the
    # software. Run the test with `brew test lazyhub`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    system "false"
  end
end
