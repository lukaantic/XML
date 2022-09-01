import React, { useContext } from 'react';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Modal from '@mui/material/Modal';
import { TextField } from '@mui/material';
import { PostContext } from '../../context';



const NewPostModal = () => {
  const { createNewPost } = useContext(PostContext);
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [image, setImage] = React.useState();
  const [previewImage, setPreviewImage] = React.useState();

  const handleChangeImage = () => (event) => {
    if (event.target.files) {
      const file = event.target.files[0];
      setImage(file);
      fileToBase64(file);
    }
  };

  const fileToBase64 = async (file) => {
    if (file !== undefined) {
      const reader = new FileReader();

      reader.onloadend = () => {
        let base64data = reader.result;
        setPreviewImage(base64data);
      };

      reader.readAsDataURL(file);
      return;
    }
    setImage(null);
  };
  const [postData, setPostData] = React.useState({
    link: '',
    description: '',
    // image
  });

  const handleChange = (name, value) => {
    setPostData({ ...postData, [name]: value });
  };

  const submitNewPost = () => {
    createNewPost(postData)
    setOpen(false);
  }

  return (
    <div>
      <Button style={styles.bottomMargin} onClick={handleOpen}>Create new post</Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={styles.boxStyle}>
          <div style={{ display: "flex", flexDirection: "column" }}>
            <img
              src={previewImage}
              style={styles.imageStyle}
            />
            <label htmlFor='contained-button-file'>
              <input
                id='contained-button-file'
                accept="image/*"
                name="myImage"
                title="Upload post image"
                style={styles.inputField}
                type="file"
                onChange={handleChangeImage()}
              />
              <button
                style={styles.inputButton}
                variant='contained'
                component='span'
              >
                <h1>Upload your post picture</h1>
              </button>
            </label>

          </div>
          <div style={styles.verticalMargin}>
            <TextField
              name="description"
              required
              fullWidth
              id="description"
              label="Description"
              autoFocus
              onChange={(event) => handleChange('description', event.target.value)}
            />
          </div>
          <TextField
            name="link"
            fullWidth
            id="link"
            label="Link"
            autoFocus
            onChange={(event) => handleChange('link', event.target.value)}
          />
          <Button
            onClick={() => submitNewPost()}
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Submit new post
          </Button>
        </Box>
      </Modal>
    </div >
  );
}

export default NewPostModal;

const styles = {
  imageStyle: {
    alignSelf: "center",
    width: "200px",
    height: "200px",
    display: "flex",
    alignItems: "center",
    overflow: "hidden",
    borderWidth: "4px",
    objectFit: "cover",
    marginBottom: 30
  },
  inputField: {
    zIndex: 1,
    position: "absolute",
    opacity: 0,
    width: "86%",
    height: "60px",
    cursor: "pointer"
  },
  inputButton: {
    position: "absolute",
    width: "86%",
    height: "60px",
    cursor: "pointer"
  },
  verticalMargin: {
    marginBottom: 30,
    marginTop: 90,
  },
  bottomMargin: {
    marginBottom: 30
  },
  boxStyle: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    border: '2px solid #000',
    boxShadow: 24,
    p: 4,
    marginBottom: 30
  }
};