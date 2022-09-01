import * as  userApiClient from '../apiClient/userApiClient';

export const getAllPublicUsers = async () => {
  const users = await userApiClient.getAllPublicUsers();
  return users;
};

export const getUsersPosts = async (username) => {
  const users = await userApiClient.getUsersPosts(username);
  return users;
};
