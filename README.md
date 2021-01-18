
This is a repository to demonstrate how to use the Wails library to host an HTML/JS web client in a native/desktop Go client.
The web client here is written in vanilla JS, and compiled with esbuild.


# Setup:
```
cd client
npm run setup
```

# To run:
```
wails build
cd build
./example
```

# For development:
```
wails serve
cd client
npm run serve
```