import axios from 'axios';

export const API_BASE_URL = 'http://127.0.0.1:8081';
const token = window.localStorage.getItem('token');

export const client = axios.create({
  baseURL: "",
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
    'token': token,
  }
});