<?php



/*请完成下面这个函数，实现题目要求的功能
当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
******************************开始写代码******************************/
function smallestRange() {
    $str = fgets(STDIN);
    $k=$str[0];
    $str = fgets(STDIN);
    $arr=explode(' ',$str);
    $max=0;
    $min=0;
    $now=0;
    $tmp=0;
    foreach($arr as $v){
        $tmpmax=$v+$k;
        $tmpmin=$v-$k;
        if($tmpmin>=$max){
            $tmpnow=$tmpmin-$max;
            if($tmpnow<$now){
                $max=$tmpmin;
                $min=$max;
                $now=$tmpnow;
            }
        }else{
            $tmpnow=$max-$tmpmin;
            if($tmpnow<$now){
                $max=$max;
                $min=$tmpmin;
                $now=$tmpnow;
            }
        }
        if ($tmpmin>=$min) {
            $tmpnow=$tmpmin-$min;
            if ($tmpnow<$now) {
                $max=$tmpmin;
                $min=$min;
                $now=$tmpnow;
            }
        }else{
            $tmpnow=$min-$tmpmin;
            if($tmpnow<$now){
                $max=$min;
                $min=$tmpmin;
                $now=$tmpnow;
            }
        }

        if($tmpmax>=$max){
            $tmpnow=$tmpmax-$max;
            if($tmpnow<$now){
                $max=$tmpmax;
                $min=$max;
                $now=$tmpnow;
            }
        }else{
            $tmpnow=$max-$tmpmax;
            if($tmpnow<$now){
                $max=$max;
                $min=$tmpmax;
                $now=$tmpnow;
            }
            if ($tmpmax>=$min) {
                $tmpnow=$tmpmax-$min;
                if ($tmpnow<$now) {
                    $max=$tmpmax;
                    $min=$min;
                    $now=$tmpnow;
                }
            }else{
                $tmpnow=$min-$tmpmax;
                if($tmpnow<$now){
                    $max=$min;
                    $min=$tmpmax;
                    $now=$tmpnow;
                }
            }
        }
    }
    return $now;
    
}
/******************************结束写代码******************************/


$__fp = fopen("php://stdin", "r");
  
$res = smallestRange();
print($res . "\n" );

fclose($__fp);

?>
