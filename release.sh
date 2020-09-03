./.set-token.sh

git tag -a v0.0.1 -m "Release"
git push origin v0.0.1
rm -rf dist/
goreleaser release --skip-validate