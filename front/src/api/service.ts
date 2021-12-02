import axios from 'axios';

export const API_BASE_URL = '';

export const client = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
    'lang': 'en_US',
  }
});

export const authClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
    'lang': 'en_US',
  }
});

authClient.interceptors.request.use(function(config){
  config.headers['Authorization']='Bearer ' + window.sessionStorage.getItem('token')
  return config
},function(error){
  return Promise.reject(error)
})

authClient.interceptors.response.use(response=>{
   return response
},function(error){
   if(error.response.status ===401){
    window.sessionStorage.setItem('token',"");
    window.sessionStorage.setItem('username',"");
    window.location.href = '/login';
   }
   return Promise.reject(error)
}
)