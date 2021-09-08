# delay-message-pluigins
Implementation delay message using RabbitMQ Plugins and Golang

### Requirement
#### Application
- RabbitMQ@3.9.3
- Go@1.16

#### Environment
Please set environment variable for setup this aplication
<!DOCTYPE html>
<html>
<body>
  
 <table>
 	<tr>
 		<td> Name</td>
 		<td> Example Value</td>
 	</tr>
 	<tr>
 		<td> Driver </td>
 		<td> amqp </td>
 	</tr>
  <tr>
 		<td> User </td>
 		<td> admin </td>
 	</tr>
  <tr>
 		<td> Password </td>
 		<td> admin </td>
 	</tr>
  <tr>
 		<td> Host </td>
 		<td> localhost </td>
 	</tr>
   <tr>
 		<td> Port </td>
 		<td> 5672 </td>
 	</tr>
 </table>

</body>
</html>

#### Plugins
Plugins are used to help the process of making this application so that some of these plugins are mandatory or must be installed

<!DOCTYPE html>
<html>
<body>
  
 <table>
 	<tr>
 		<td> Name </td>
 		<td> Owner </td>
    <td> Data </td>
    <td> Doc </td>
 	</tr>
 	<tr>
 		<td> rabbitmq_delayed_message_exchange </td>
 		<td> RabbitMQ </td>
    <td> https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases </td>
    <td> https://www.rabbitmq.com/installing-plugins.html </td>
 	</tr>
 </table>

</body>
</html>

#### Run Program
1) Main Execution

      ```bash
      go run app/main.go
      ```
2) if you want to put the env file in another folder not in the main folder then run the program with this

      ```bash
      ENV = "lala/.env" go run app/main.go
      ```
 
