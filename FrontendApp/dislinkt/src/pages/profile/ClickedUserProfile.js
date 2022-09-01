import React, { useState, useContext, useEffect, useMemo } from 'react';
import UserPost from '../../components/post/UsersPost';
import Header from '../../components/navigation/Header';
import { UserContext } from '../../context';
import UserAccountForDisplay from '../../components/users/UserAccountForDisplay';
import { useParams } from 'react-router-dom';


const ClickedUserProfile = () => {
  let { id } = useParams();

  const { getUsersPosts, usersPosts } = useContext(UserContext);

  useEffect(() => {
    if (usersPosts == 0) {
      getUsersPosts(id);

    }
    console.log('usersPosts', usersPosts)
  }, [usersPosts])


  return (
    <div>
      <Header />
      <div style={{ display: "flex", flex: 1, flexDirection: "column", alignItems: "center", width: "100%", margin: 50 }}>
        {usersPosts?.length > 0 &&
          usersPosts?.map((publicProfile, index) => {
            return (
              <UserPost key={index} comment={publicProfile?.Comment} username={publicProfile?.RegularUser.Username} description={publicProfile?.Description} dislikes={publicProfile?.Dislikes} likes={publicProfile?.Likes} link={publicProfile.link} />
            )
          })
        }
      </div>
    </div>
  );
}

export default ClickedUserProfile;
