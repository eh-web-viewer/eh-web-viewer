# 他妈到底怎么新建vue啊

每次都忘

## init

```sh
npm create vite@latest .

```
选择 vue，TS

## proxy

https://vitejs.dev/config/server-options.html#server-proxy

## path

tsconfig.json
```json
{
  "compilerOptions": {    
    "paths": { "@/*":["./src/*"] },
  }
}
```

vite.config.ts
```ts
import path from 'path'
  resolve: {
    alias:{
      '@': path.resolve(__dirname, './src')
    }
  }
```

cant find __dirname and path

    npm install @types/node -D

but it seems not working

## router


### dynamic router


get route in setup 

https://stackoverflow.com/questions/65284428/how-to-get-params-of-router-in-vue-3