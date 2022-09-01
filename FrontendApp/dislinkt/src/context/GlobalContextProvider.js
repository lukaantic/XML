import React from 'react';

import { AuthenticationContextProvider } from './AuthenticationContext';
import { PostContextProvider } from './PostContext';
import { UserContextProvider } from './UserContext';

const GlobalContextProvider = props => {
  return (
    <UserContextProvider>
      <AuthenticationContextProvider>
        <PostContextProvider>
          {props.children}
        </PostContextProvider>
      </AuthenticationContextProvider>
    </UserContextProvider>
  );
};

export default GlobalContextProvider;
