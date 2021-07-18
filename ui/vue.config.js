module.exports = {
  chainWebpack: config => {
    config
      .plugin('html')
      .tap(args => {
        args[0].title = 'Short{Paste}'
        return args
      })
  },
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
