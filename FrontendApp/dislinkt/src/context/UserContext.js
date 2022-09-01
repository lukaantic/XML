import React, { createContext, useState, useEffect, } from 'react';

import { userService } from '../services';
// import { UserContext } from '.';

const UserContext = createContext();

const UserContextProvider = props => {
  const [publicProfiles, setPublicProfiles] = useState([]);
  const [usersPosts, setUsersPosts] = useState([]);

  async function getAllPublicUsers() {
    try {
      const response = await userService.getAllPublicUsers();

      setPublicProfiles(response?.data);

      return response?.data;

    } catch (error) {
      console.log('Error while getting public accounts :>> ', error?.message);
    }
  }

  async function getUsersPosts(username) {
    try {
      const response = await userService.getUsersPosts(username);

      setUsersPosts(response?.data);

      return response?.data;

    } catch (error) {
      console.log('Error while getting public accounts :>> ', error?.message);
    }
  }


  const providerValue = {
    getAllPublicUsers,
    getUsersPosts,
    usersPosts,
    publicProfiles
  };

  return (
    <UserContext.Provider value={providerValue}>
      {props.children}
    </UserContext.Provider>
  );
};

export { UserContext, UserContextProvider };
