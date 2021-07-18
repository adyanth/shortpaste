module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    proxy: {
      "/api/*": {
        target: "http://localhost:8000",
        secure: false
      },
      "/l/*": {
        target: "http://localhost:8000",
        secure: false
      },
      "/t/*": {
        target: "http://localhost:8000",
        secure: false
      },
      "/f/*": {
        target: "http://localhost:8000",
        secure: false
      },
    }
  }
}
