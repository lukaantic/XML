import { axiosInstance } from '../axios/axios';

export const login = async payload => {
  const data = await axiosInstance.post('auth/login', payload);
  return data;
};

export const signUp = async payload => {
  const data = await axiosInstance.post('/user/register', payload);
  return data;
};