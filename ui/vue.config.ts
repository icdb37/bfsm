import { defineConfig } from '@vue/cli-service';
import path from 'path';

export default defineConfig({
  transpileDependencies: ['@dcloudio/uni-ui'],
  chainWebpack: (config) => {
    config.resolve.alias
      .set('@', path.resolve(__dirname))
      .set('xapi', path.resolve(__dirname, 'xapi'))
  }
})