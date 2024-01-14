const env = import.meta.env.MODE || 'prod'

const EnvConfig = {

    dev:{
        baseApi: '/api',
        mockApi: 'http://localhost:5173'
    },
    test:{
        baseApi: '/api',
        mockApi: 'http://localhost:5173'
    },
    prod:{
        baseApi: '/api',
        mockApi: 'http://localhost:5173'
    },
}

export default {
    env,
    mock: true,
    ...EnvConfig[env]
}