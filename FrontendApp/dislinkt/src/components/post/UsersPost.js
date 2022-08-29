import React from "react";

const urls = ["Url numero uno", "url numero duo"];

const UserPost = ({}) => {
    return (
        <div onClick={console.log("Post pressed")}>
            <div style={{flexDirection: 'row', display: 'flex'}}>
                    <img src="https://reactjs.org/logo-og.png" alt="React Logo"  style={styles.postImage} />
                <div>
                    
                </div>
                <div style={styles.mainContentRowContainer}>
                    Some titile bla bla
                </div>
            </div>
            <div>
                {urls.map(( url, index ) => {
                    <div>
                        {url[index]}
                    </div>
                })}
            </div>
        </div>
    );    
}

const styles = {
    postImage: { width: 60, height: 60, borderRadius: 5 },
    mainContentRowContainer: { flex: 1, flexDirection: 'row', height: 60, display: 'flex' },
    locationNtileRowContainer: {
      justifyContent: 'space-between',
      flexDirection: 'row',
      paddingTop: 6,
      flex: 1,
      display: 'flex'
    },
    midContentContainer: {
      flexDirection: 'row',
      justifyContent: 'space-between',
      marginTop: 10,
    },
    sellerContentContainer: {
      flexDirection: 'row',
      justifyContent: 'space-between',
      alignItems: 'center',
    },
    sellerNameTextStyle: { paddingLeft: 8, fontWeight: 'bold' },
    rowCenterAligned: { flexDirection: 'row', alignItems: 'center' },
    rateTextStyle: { marginHorizontal: 5, fontSize: 12 },
  }

  export default UserPost;