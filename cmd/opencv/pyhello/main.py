import cv2 as cv
img = cv.imread("/home/siuyin/Downloads/old/face-20240423-120243.jpg")

cv.imshow("Display window", img)
k = cv.waitKey(0) 