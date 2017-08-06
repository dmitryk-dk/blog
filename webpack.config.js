const path = require('path');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
const WebpackNotifierPlugin = require('webpack-notifier');
const ProgressBarPlugin = require('progress-bar-webpack-plugin');

module.exports = {
    entry: path.resolve(__dirname, 'frontend/index.js'),
    output: {
        path: path.resolve(__dirname, './dist'),
        filename: 'bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                loader: 'babel-loader',
                query: {
                    presets: ['babel-preset-react', 'babel-preset-es2015', 'babel-preset-stage-2'].map(require.resolve),
                    cacheDirectory: true
                }
            },
            {
                test: /\.scss$/,
                use: ExtractTextPlugin.extract({
                    fallback: 'style-loader',
                    use: [
                        {
                            loader: 'css-loader',
                        },
                        {
                            loader: "sass-loader"
                        },
                    ]
                })
            }
        ]
    },
    resolve: {
        modules: [
            path.resolve(__dirname, 'node_modules'),
        ],
        extensions: ['.js', '.jsx', '.scss', '.json'],
    },
    plugins: [
        new WebpackNotifierPlugin({
            alwaysNotify: true
        }),
        new ProgressBarPlugin(),
        new ExtractTextPlugin('app.css')
    ],
    devtool: false,
    stats: {
        children: false
    }
};
