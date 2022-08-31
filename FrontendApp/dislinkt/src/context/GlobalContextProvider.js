import React from 'react';

import { AuthenticationContextProvider } from './AuthenticationContext';
// import {UserContextProvider} from './UserContext';
// import {TrainerContextProvider} from './TrainerContext';
// import {LocationContextProvider} from './LocationContext';
// import {ExerciserContextProvider} from './ExerciserContext';

const GlobalContextProvider = props => {
  return (
    // <UserContextProvider>
    <AuthenticationContextProvider>
      {/* <LocationContextProvider> */}
      {/* <TrainerContextProvider> */}
      {/* <ExerciserContextProvider> */}
      {props.children}
      {/* </ExerciserContextProvider> */}
      {/* </TrainerContextProvider> */}
      {/* </LocationContextProvider> */}
    </AuthenticationContextProvider>
    // </UserContextProvider>
  );
};

export default GlobalContextProvider;
