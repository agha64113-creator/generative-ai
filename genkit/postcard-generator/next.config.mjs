/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  // Workaround Handlebar import issues
  // https://github.com/handlebars-lang/handlebars.js/issues/953
  webpack: (config) => {
    config.resolve.alias["handlebars"] = "handlebars/dist/handlebars.js";
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
  },
  // Performance optimizations
  compress: true,
  poweredByHeader: false,
  // Image optimization
  images: {
    formats: ["image/avif", "image/webp"],
    deviceSizes: [640, 750, 828, 1080, 1200, 1920, 2048, 3840],
    imageSizes: [16, 32, 48, 64, 96, 128, 256, 384],
    minimumCacheTTL: 60,
  },
  // Bundle optimization
  swcMinify: true,
  compiler: {
    removeConsole: process.env.NODE_ENV === "production" ? {
      exclude: ["error", "warn"],
    } : false,
  },
  // Output optimization
  output: "standalone",
};

export default nextConfig;
