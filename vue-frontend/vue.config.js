const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  assetsDir: "static",
  devServer: {
    proxy:{
      '^/api':{
        target: 'http://192.168.211.129:5000',
        changeOrigin: true,
        logLevel:'debug',
        pathRewrite: {'^/api': '/'}
      }
    }
  }
}
)
