import React, { useState, useContext, useEffect, useMemo } from 'react';
import UserPost from '../../components/post/UsersPost';
import Header from '../../components/navigation/Header';
import { UserContext } from '../../context';
import UserAccountForDisplay from '../../components/users/UserAccountForDisplay';
import NewPostModal from '../../components/post/NewPostModal';


const Home = () => {
  const { getAllPublicUsers, publicProfiles } = useContext(UserContext);

  useEffect(() => {
    if (publicProfiles == 0) {
      getAllPublicUsers();
    }
  }, [publicProfiles])

  return (
    <div>
      <Header />
      <div style={{ display: "flex", flex: 1, flexDirection: "column", alignItems: "center", width: "100%", margin: 50 }}>
        <NewPostModal />
        {publicProfiles.length > 0 &&
          publicProfiles.map((publicProfile, index) => {
            return (
              <UserAccountForDisplay key={index} username={publicProfile?.username} name={publicProfile?.name} surname={publicProfile?.surname} email={publicProfile?.email} />
            )
          })
        }
      </div>
    </div>
  );
}

export default Home;
