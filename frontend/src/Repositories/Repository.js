import axios from 'axios'

const baseDomain = 'localhost:8080'
const baseURL = `http://${baseDomain}/api`
const token = localStorage.getItem("token")

export default axios.create({
    baseURL: baseURL,
    headers: {
        'Authorization': `Bearer ${token}`
    }
})