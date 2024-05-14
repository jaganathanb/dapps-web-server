# DApps Web Server

#Create tag & push to remote

```

git tag <version number> # ex. git tag v1.4.1.0

git push origin v1.4.0.0 # pushing the tag to remote

```

#Build

```
go-winres make --product-version=git-tag --file-version=git-tag

#open bash terminal and run the below cmd

bash build.sh

```
