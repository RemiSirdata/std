<?php $errorMSG = "";

// NAME
if (empty($_POST["name"])) {
    $errorMSG = "Name is required ";
} else {
    $name = $_POST["name"];
}

// EMAIL
if (empty($_POST["email"])) {
    $errorMSG .= "Email is required ";
} else {
    $email = $_POST["email"];
}

// Number of Guest
if (empty($_POST["guest"])) {
	$errorMSG .= "Field is required ";
} else {
    $guest = $_POST["guest"];
}


// Attending Event
if (empty($_POST["evente"])) {
    $errorMSG .= "Event is required ";
} else {
    $event = $_POST["evente"];
}


// MESSAGE
if (empty($_POST["message"])) {
    $errorMSG .= "Message is required ";
} else {
    $message = $_POST["message"];
}


$EmailTo = "yourname@domain.com";
$Subject = "Subject";

// prepare email body text
$Body = "";
$Body .= "Name: ";
$Body .= $name;
$Body .= "\n";
$Body .= "Email: ";
$Body .= $email;
$Body .= "\n";
$Body .= "Number of Guest: ";
$Body .= $guest;
$Body .= "\n";
$Body .= "Attending Event: ";
$Body .= $event;
$Body .= "\n";
$Body .= "Message: ";
$Body .= $message;
$Body .= "\n";

// send email
$success = mail($EmailTo, $Subject, $Body, "From:".$email);

// redirect to success page
if ($success && $errorMSG == ""){
   echo "success";
}else{
    if($errorMSG == ""){
        echo "The message cannot be sent right now";
    } else {
        echo $errorMSG;
    }
}

?>
