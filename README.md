# Heroku Buildpack: Revel

This is a [Heroku buildpack](http://devcenter.heroku.com/articles/buildpacks) for [Revel](http://robfig.github.com/revel/).

This repository is useful if you want to inspect or change the behavior of the buildpack itself.

## Hacking

To change this buildpack, fork it on GitHub. Push
changes to your fork, then create a test app with
`--buildpack YOUR_GITHUB_GIT_URL` and push to it. If you
already have an existing app you may use `heroku config
add BUILDPACK_URL=YOUR_GITHUB_GIT_URL` instead of
`--buildpack`.
