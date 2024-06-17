import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import ToDoForm from './components/ToDoForm'
import ToDoList from './components/ToDoList'
function App() {

  return (
    <Stack h='100vh'>
      <Navbar />
      <Container>
        <ToDoForm/>
        <ToDoList/> 
      </Container>
    </Stack>
  )
}

export default App
