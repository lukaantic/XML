import React, { createContext, useState, useEffect, } from 'react';

import { postService } from '../services';

const PostContext = createContext();

const PostContextProvider = props => {

  async function createNewPost(data, successCallback, errorCallback) {
    try {
      const payload = {
        username: localStorage.getItem('username'),
        description: data.description,
        mediaPath: data.imagePath,
        link: data.link
      }

      const response = await postService.createNewPost(payload);

      console.log(response?.statusText);
    } catch (error) {
      alert("Error while creating new post: " + error.message);
    }
  }

  const providerValue = {
    createNewPost
  };

  return (
    <PostContext.Provider value={providerValue}>
      {props.children}
    </PostContext.Provider>
  );
};

export { PostContext, PostContextProvider };
