import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '');
  const portValue = env.SPA_PORT;

  if (!portValue) {
    throw new Error('SPA_PORT environment variable is required');
  }

  const port = Number(portValue);
  if (Number.isNaN(port)) {
    throw new Error('SPA_PORT must be a valid number');
  }

  return {
    plugins: [vue()],
    server: {
      port,
    },
    build: {
      target: 'esnext',
    },
    test: {
      environment: 'happy-dom',
      globals: true,
      coverage: {
        provider: 'v8',
        include: [
          'src/application/**/*.{ts,vue}',
          'src/composables/**/*.{ts,vue}',
          'src/infrastructure/**/*.{ts,vue}',
        ],
        exclude: ['src/assets/**', 'src/components/**', 'src/application/SimulationState.ts'],
      },
    },
  };
});
