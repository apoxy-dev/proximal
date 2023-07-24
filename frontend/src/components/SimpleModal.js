import Box from '@mui/material/Box';
import styled from '@mui/material/styles/styled';

const SimpleModal = styled(Box)({
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  minWidth: 400,
  background: '#FFF',
  border: '1px solid #000',
  borderRadius: '3px',
  padding: '2rem',
});

export default SimpleModal;
