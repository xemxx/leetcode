<?php



/*请完成下面这个函数，实现题目要求的功能
当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^
******************************开始写代码******************************/
function find_longest_num_str($input)
{
    $count=0;
    $maxCount=0;
    $output='';
    $maxOutput='';
    for ($i=0;$i<strlen($input);$i++) {
        if(is_numeric($input[$i])){
            $count++;
            $output=$output.$input[$i];
        }else{
            if($maxCount<$count){
                $maxCount=$count;
                $maxOutput=$output;
            }
            $count=0;
            $output='';
        }
    }
    if($maxCount<$count){
        $maxCount=$count;
        $maxOutput=$output;
    }
    $count=0;
    $output='';
    return "$maxCount/$maxOutput";
}
/******************************结束写代码******************************/


$__fp = fopen("php://stdin", "r");

$_input = fgets($__fp);
$_input = trim($_input);

  
$res = find_longest_num_str($_input);
echo $res;
//fwrite($res . "\n",1);

fclose($__fp);
