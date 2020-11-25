<?php
// Ran this in the interactive shell packaged in the PHP download. Unsure of the standard way things are ran in PHP.
function getAge() {
    $age;
    $isNum = false;
    do {
        $age = readline("Enter your age: ");

        if (is_numeric($age))
            $isNum = true;
        else
            print("Age must be a number.\n");
    } while (!$isNum);
    print("Your age is ". $age . ".");
}

getAge();
?>
