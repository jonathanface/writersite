/* eslint-disable */
import { defineConfig } from 'vitest/config';
import react from '@vitejs/plugin-react';
import type { Plugin } from 'vite';

// Proxy single-segment paths (e.g. /thecow) to the Go backend
// so download short URLs work in dev mode
function downloadShortUrlProxy(): Plugin {
    return {
        name: 'download-short-url-proxy',
        configureServer(server) {
            server.middlewares.use((req, res, next) => {
                const path = req.url || '';
                // Only match single-segment paths like /thecow (no dots, no slashes after the first)
                if (/^\/[a-zA-Z0-9_-]+$/.test(path)) {
                    const target = `http://localhost:8443${path}`;
                    import('http').then((http) => {
                        http.get(target, (proxyRes) => {
                            if (proxyRes.statusCode === 301 || proxyRes.statusCode === 302 || proxyRes.statusCode === 307) {
                                res.writeHead(proxyRes.statusCode, proxyRes.headers);
                                res.end();
                            } else {
                                // Not a download slug — let Vite handle it
                                next();
                            }
                        }).on('error', () => {
                            next();
                        });
                    });
                } else {
                    next();
                }
            });
        },
    };
}

// https://vite.dev/config/
export default defineConfig({
    plugins: [react(), downloadShortUrlProxy()],
    define: {
        global: {},
    },
    server: {
        port: 8080,
        cors: true,
        proxy: {
            '/api': {
                target: 'http://localhost:8443',
                changeOrigin: true,
                secure: false,
                configure: (proxy, _options) => {
                    proxy.on('error', (err, _req, _res) => {
                        console.error('proxy error', err);
                    });
                    proxy.on('proxyReq', (_proxyReq, req, _res) => {
                        console.info('Sending Request to the Target:', req.method, req.url);
                    });
                    proxy.on('proxyRes', (proxyRes, req, _res) => {
                        console.log('Received Response from the Target:', proxyRes.statusCode, req.url);
                    });
                },
            },
            '/downloads': {
                target: 'http://localhost:8443',
                changeOrigin: true,
                secure: false,
            },
        },
    },
    test: {
        globals: true, // Use Jest-like globals (describe, it, expect)
        environment: 'jsdom', // Simulate browser-like environment
        setupFiles: './vitest.setup.ts', // Optional setup file for global configurations
    },
});