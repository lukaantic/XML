import * as  postApiClient from '../apiClient/postApiClient';

export const createNewPost = async (data) => {
  const users = await postApiClient.createNewPost(data);
  return users;
};

