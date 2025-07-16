/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
    images: {
        domains: ["cdn.corpchat.com", "avatars.dicebear.com"],
          },
          };
          
module.exports = nextConfig;