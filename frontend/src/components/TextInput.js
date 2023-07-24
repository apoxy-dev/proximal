import InputBase from '@mui/material/InputBase';
import { styled } from '@mui/material/styles';

const TextInput = styled(InputBase)(({ theme }) => ({
  border: '1px solid',
  borderColor: theme.palette.text.secondary,
  padding: '8px 16px',
  margin: '8px 0',
  borderRadius: '3px',
}));

export default TextInput;
