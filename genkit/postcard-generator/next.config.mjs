/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true, // Use SWC for faster minification
  compress: true, // Enable gzip compression
  
  // Optimize production output
  output: 'standalone',
  
  // Image optimization
  images: {
    formats: ['image/avif', 'image/webp'],
    minimumCacheTTL: 60,
  },
  
  // Workaround Handlebar import issues
  // https://github.com/handlebars-lang/handlebars.js/issues/953
  webpack: (config, { dev, isServer }) => {
    config.resolve.alias["handlebars"] = "handlebars/dist/handlebars.js";
    
    // Production optimizations
    if (!dev && !isServer) {
      config.optimization = {
        ...config.optimization,
        moduleIds: 'deterministic',
        splitChunks: {
          chunks: 'all',
          cacheGroups: {
            default: false,
            vendors: false,
            vendor: {
              name: 'vendor',
              chunks: 'all',
              test: /node_modules/,
              priority: 20,
            },
            common: {
              name: 'common',
              minChunks: 2,
              chunks: 'all',
              priority: 10,
              reuseExistingChunk: true,
              enforce: true,
            },
            mui: {
              name: 'mui',
              test: /@mui/,
              chunks: 'all',
              priority: 30,
            },
            genkit: {
              name: 'genkit',
              test: /@genkit-ai/,
              chunks: 'all',
              priority: 30,
            },
          },
        },
      };
    }
    
    return config;
  },
  // Workaround dependency errors
  // https://github.com/open-telemetry/opentelemetry-js/pull/4214
  experimental: {
    instrumentationHook: true,
    serverComponentsExternalPackages: [
      "@opentelemetry/auto-instrumentations-node",
      "@opentelemetry/sdk-node",
    ],
    optimizeCss: true, // Enable CSS optimization
    optimizePackageImports: ['@mui/material', '@mui/icons-material', '@mui/lab'], // Tree-shake Material-UI
  },
};

export default nextConfig;
