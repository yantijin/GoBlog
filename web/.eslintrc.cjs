module.exports = {
  // 检测环境 ，浏览器、es、node
  env: {
    browser: true,
    es2021: true,
    node: true,
  },

  // 用了哪些别人定义好的规则
  extends: ['plugin:vue/vue3-strongly-recommended', 'prettier'], // 'plugin:import/typescript',
  // "extends": [ "eslint:recommended","plugin:vue/essential","plugin:@typescript-eslint/recommended"],
  // ESLint 会对我们的代码进行校验,而parser的作用是将我们写的代码转换为ESTree(AST),ESLint 会对ESTree进行校验
  parser: 'vue-eslint-parser',
  // 解析器配置
  parserOptions: {
    // es的版本号,或者年份都可以,latest代表"最近的"
    ecmaVersion: 'latest',
    parser: '@typescript-eslint/parser',
    // 源码类型,默认是script,es模块使用module
    sourceType: 'module',
    // 额外的语言类型
    ecmaFeatures: {
      tsx: true,
      jsx: true,
    },
  },
  // 全局自定义的宏,这样在源文件中使用全局变量就不会报错或者警告
  globals: {
    defineProps: 'readonly',
    defineEmits: 'readonly',
    defineExpose: 'readonly',
    withDefaults: 'readonly',
  },

  // 插件
  // 前缀eslint-plugin- 可以省略
  // vue官方提供了一个ESLint插件 eslint-plugin-vue,它提供了 parser 和 rules
  // parser为vue-eslint-parser,放在上面的parser字段里
  // rules放在extends字段里,选择合适的规则
  plugins: [
    'vue',
    'eslint-plugin-import',
    // "@typescript-eslint"
  ],
  settings: {
    'import/resolver': {
      alias: {
        map: [['@', './src']],
      },
    },
    // 允许的扩展名
    'import/extensions': ['.js', '.jsx', '.ts', '.tsx', '.mjs'],
  },
  // 自定义规则,覆盖上面extends继承的第三方规则,根据自己项目灵活定义
  rules: {
    'import/no-extraneous-dependencies': 0,
    'no-param-reassign': 0,
    'vue/multi-word-component-names': 0,
    'vue/attribute-hyphenation': 0,
    'vue/v-on-event-hyphenation': 0,
  },
};
