import com.ceacar.ClockAngel;
//package com.ceacar;
public class Main { 
	public static void main(String [] args)
	{
	  int h = 12;
	  int m = 30;
	  ClockAngel ca = new ClockAngel();
	  int angle = ca.getAngel(h, m);
	  System.out.println(angle);
	}

}
