import React, { createContext, useState, useEffect, } from 'react';

import { axiosInstance } from '../axios/axios';

import { authenticationService } from '../services';
// import { UserContext } from '.';

const AuthenticationContext = createContext();

const AuthenticationContextProvider = props => {
  const [token, setToken] = useState('');

  const setUserToken = token => {
    setToken(token);
    axiosInstance.defaults.headers.authorization = `Bearer ${token}`;
  };

  const saveUserToken = token => {
    setUserToken(token);
    return localStorage.setItem('@token', token);
  };

  const deleteUserToken = () => {
    axiosInstance.defaults.headers.authorization = '';
    setToken('');
    // deleteUserData();
    return localStorage.removeItem('@token');
  };

  const getUserToken = () => {
    localStorage.getItem('@token', (error, userToken) => {
      if (userToken)
        return setUserToken(userToken);

      setUserToken(null)
    });
  };

  useEffect(() => {
    if (token === '')
      getUserToken();
  }, [token]);

  async function login(data, successCallback, errorCallback) {
    try {
      const response = await authenticationService.login(data);

      saveUserToken(response.data);

      successCallback();
    } catch (error) {
      errorCallback(error?.message);
    }
  }
  async function signUp(data, successCallback, errorCallback) {
    try {
      const response = await authenticationService.signUp(data);

      saveUserToken(response.data.token);

      successCallback(response?.statusText);
    } catch (error) {
      errorCallback("Error while signing up, probably account already exists");
    }
  }

  const providerValue = {
    token,
    deleteUserToken,
    login,
    signUp,
  };

  return (
    <AuthenticationContext.Provider value={providerValue}>
      {props.children}
    </AuthenticationContext.Provider>
  );
};

const LoginConsumer = AuthenticationContext.Consumer;

export { AuthenticationContext, AuthenticationContextProvider, LoginConsumer };
