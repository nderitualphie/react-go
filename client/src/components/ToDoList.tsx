import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import {useQuery} from "@tanstack/react-query"
import ToDoItem from "./ToDoItem";
import { BASE_URL } from "../App";
export type Todo = {
	_id: number;
	body: string;
	completed: boolean;
};
const ToDoList = () => {
	
	const {data:todos,isLoading}=useQuery<Todo[]>({
		queryKey:["todos"],
		queryFn:async ()=> {
			try {
				const res = await fetch(BASE_URL + "/todo")
				const data = await res.json()
				if(!res.ok){
					throw new Error(data.error ||"something went wrong");
					
				}
				return data || []
			} catch (error) {
				console.log(error)
			}
		}
	})
	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}
             bgGradient='linear(to-l, #0b85f8, #00ffff)' bgClip='text'>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading && todos?.length === 0 && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			<Stack gap={3}>
				{todos?.map((todo) => (
					<ToDoItem key={todo._id} todo={todo} />
				))}
			</Stack>
		</>
	);
};
export default ToDoList;