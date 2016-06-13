<?php

$file = file_get_contents('axe.html');
$doc = new DOMdocument; 
$doc -> loadHTML ($file);

$tr_doms = $doc -> getElementsByTagName('tr');
 $array = array();
  foreach ($tr_doms as $tr_dom) {
    $td_doms = $tr_dom->getElementsByTagName('td');
      $name = $td_doms->item(0)->nodeValue;
      $chinese = intval($td_doms->item(1)->nodeValue);
      $math = intval($td_doms->item(2)->nodeValue);
      $sci = intval($td_doms->item(3)->nodeValue);
      $social = intval($td_doms->item(4)->nodeValue);
      $health = intval($td_doms->item(5)->nodeValue);
  
    if ($name == '姓名') continue;
    
    $array_grade = array('國語'=> $chinese, '數學'=> $math, 
                         '自然'=> $sci, '社會'=> $social, 
                         '健康教育'=> $health);
     
    $array[] = array('name'=> $name, 'grades'=> $array_grade);             
      
   }

echo json_encode($array, JSON_UNESCAPED_UNICODE);
//print_r($array); 
