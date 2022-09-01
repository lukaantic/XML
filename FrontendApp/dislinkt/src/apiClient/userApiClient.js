import { axiosInstance } from '../axios/axios';

export const login = async payload => {
  const data = await axiosInstance.post('auth/login', payload);
  return data;
};

export const signUp = async payload => {
  const data = await axiosInstance.post('/user/register', payload);
  return data;
};

export const getAllPublicUsers = async () => {
  const data = await axiosInstance.get('/user/public-regular-users');
  return data;
}

export const getUsersPosts = async (username) => {
  const data = await axiosInstance.get(`/post/regular-user-posts/${username}`);
  return data;
}