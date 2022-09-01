import { axiosInstance } from '../axios/axios';

export const createNewPost = async (data) => {
  const response = await axiosInstance.post("/post/new-post", data);
  return response;
}