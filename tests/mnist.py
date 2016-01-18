# taken from: http://pjreddie.com/projects/mnist-in-csv/

# convert reads the binary data and outputs a csv file
def convert(image_file_path, label_file_path, csv_file_path, n):
    # open files
    images_file = open(image_file_path, "rb")
    labels_file = open(label_file_path, "rb")
    csv_file = open(csv_file_path, "w")

    # read some kind of header
    images_file.read(16)
    labels_file.read(8)

    # prepare array
    images = []

    # read images and labels
    for _ in range(n):
        image = []
        for _ in range(28*28):
            image.append(ord(images_file.read(1)))
        image.append(ord(labels_file.read(1)))
        images.append(image)

    # write csv rows
    for image in images:
        csv_file.write(",".join(str(pix) for pix in image)+"\n")

    # close files
    images_file.close()
    csv_file.close()
    labels_file.close()

# convert train data set
convert(
    "train-images-idx3-ubyte",
    "train-labels-idx1-ubyte",
    "train.csv",
    60000
)

# convert test data set
convert(
    "t10k-images-idx3-ubyte",
    "t10k-labels-idx1-ubyte",
    "test.csv",
    10000
)
