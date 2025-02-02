const argEnvIndex = process.argv.indexOf('--env')
let argEnv = (argEnvIndex !== -1 && process.argv[argEnvIndex + 1]) || ''

const RUN_ENV_MAP = {
  local: {
    instances: 1,
    max_memory_restart: '500M'
  },
  dev: {
    instances: 2,
    max_memory_restart: '500M'
  },
  prod: {
    instances: 1,
    max_memory_restart: '1G'
  }
}

if (!(argEnv in RUN_ENV_MAP)) {
  argEnv = 'prod'
}

module.exports = {
  apps: [
    {
      name: "roll-finder-api",          // Name of your app
      script: "./go-service",            // Path to your compiled Go binary
      watch: false,               // Optionally enable watch mode
      env: {                      // Default environment variables
        PORT: 8090,
        ENV: "production",
      },
      env_production: {           // Environment variables for production
        NODE_ENV: "production",
      },
    },
    // {
    //   name: 'go-service',
    //   script: './go-service', // Replace with the compiled binary of your Go service
    //   args: '', // Add additional CLI args if needed
    //   instances: RUN_ENV_MAP[argEnv].instances,
    //   exec_mode: 'cluster', // Ensure proper load balancing
    //   watch: false, // Disable watch for production
    //   max_memory_restart: RUN_ENV_MAP[argEnv].max_memory_restart,
    //   env_local: {
    //     APP_ENV: 'local',
    //     PORT: 3000, // Example environment variable
    //   },
    //   env_dev: {
    //     APP_ENV: 'dev',
    //     PORT: 4000, // Example environment variable
    //   },
    //   env_prod: {
    //     APP_ENV: 'prod',
    //     PORT: 5000, // Example environment variable
    //   }
    // }
  ]
}