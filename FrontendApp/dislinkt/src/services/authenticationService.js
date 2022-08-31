import * as  userApiClient from '../apiClient/userApiClient';

export const login = async data => {

  const payload = {
    username: data.username,
    password: data.password,
  };

  const token = await userApiClient.login(payload);
  return token;
};

export const signUp = async data => {
  const token = await userApiClient.signUp(data);
  return token;
};
